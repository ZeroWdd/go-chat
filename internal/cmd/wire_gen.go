// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"go-chat/config"
	"go-chat/internal/cmd/internal/command"
	cron2 "go-chat/internal/cmd/internal/command/cron"
	other2 "go-chat/internal/cmd/internal/command/other"
	"go-chat/internal/cmd/internal/command/queue"
	"go-chat/internal/cmd/internal/handle/cron"
	"go-chat/internal/cmd/internal/handle/other"
	queue2 "go-chat/internal/cmd/internal/handle/queue"
	"go-chat/internal/pkg/filesystem"
	"go-chat/internal/provider"
	"go-chat/internal/repository/cache"
)

// Injectors from wire.go:

func Initialize(ctx context.Context, conf *config.Config) *AppProvider {
	client := provider.NewRedisClient(conf)
	serverStorage := cache.NewSidStorage(client)
	clearWsCache := cron.NewClearWsCache(serverStorage)
	db := provider.NewMySQLClient(conf)
	filesystemFilesystem := filesystem.NewFilesystem(conf)
	clearArticle := cron.NewClearArticle(db, filesystemFilesystem)
	clearTmpFile := cron.NewClearTmpFile(db, filesystemFilesystem)
	clearExpireServer := cron.NewClearExpireServer(serverStorage)
	subcommands := &cron2.Subcommands{
		ClearWsCache:      clearWsCache,
		ClearArticle:      clearArticle,
		ClearTmpFile:      clearTmpFile,
		ClearExpireServer: clearExpireServer,
	}
	cronCommand := cron2.NewCrontabCommand(subcommands)
	queueSubcommands := &queue.Subcommands{}
	queueCommand := queue.NewQueueCommand(queueSubcommands)
	exampleHandle := other.NewExampleHandle(db)
	exampleCommand := other2.NewExampleCommand(exampleHandle)
	migrateCommand := other2.NewMigrateCommand(db)
	otherSubcommands := &other2.Subcommands{
		ExampleCommand: exampleCommand,
		MigrateCommand: migrateCommand,
	}
	otherCommand := other2.NewOtherCommand(otherSubcommands)
	commands := &command.Commands{
		CrontabCommand: cronCommand,
		QueueCommand:   queueCommand,
		OtherCommand:   otherCommand,
	}
	appProvider := &AppProvider{
		Config:   conf,
		Commands: commands,
	}
	return appProvider
}

// wire.go:

var providerSet = wire.NewSet(provider.NewMySQLClient, provider.NewRedisClient, provider.NewHttpClient, provider.NewEmailClient, provider.NewRequestClient, filesystem.NewFilesystem, cache.NewSidStorage, cron2.NewCrontabCommand, cron.NewClearTmpFile, cron.NewClearArticle, cron.NewClearWsCache, cron.NewClearExpireServer, wire.Struct(new(cron2.Subcommands), "*"), queue.NewQueueCommand, wire.Struct(new(queue.Subcommands), "*"), queue2.NewEmailHandle, other2.NewOtherCommand, other2.NewExampleCommand, other2.NewMigrateCommand, wire.Struct(new(other2.Subcommands), "*"), other.NewExampleHandle, wire.Struct(new(command.Commands), "*"), wire.Struct(new(AppProvider), "*"))
