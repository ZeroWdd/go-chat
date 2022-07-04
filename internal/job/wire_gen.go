// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"go-chat/config"
	"go-chat/internal/job/internal/command"
	cron2 "go-chat/internal/job/internal/command/cron"
	other2 "go-chat/internal/job/internal/command/other"
	"go-chat/internal/job/internal/command/queue"
	"go-chat/internal/job/internal/handle/cron"
	"go-chat/internal/job/internal/handle/other"
	queue2 "go-chat/internal/job/internal/handle/queue"
	"go-chat/internal/pkg/filesystem"
	"go-chat/internal/provider"
	"go-chat/internal/repository/cache"
	"go-chat/internal/repository/dao"
)

// Injectors from wire.go:

func Initialize(ctx context.Context, conf *config.Config) *AppProvider {
	client := provider.NewRedisClient(ctx, conf)
	sidServer := cache.NewSid(client)
	clearWsCache := cron.NewClearWsCache(sidServer)
	db := provider.NewMySQLClient(conf)
	filesystemFilesystem := filesystem.NewFilesystem(conf)
	clearArticle := cron.NewClearArticle(db, filesystemFilesystem)
	clearTmpFile := cron.NewClearTmpFile(db, filesystemFilesystem)
	clearExpireServer := cron.NewClearExpireServer(sidServer)
	subcommands := &cron2.Subcommands{
		ClearWsCache:      clearWsCache,
		ClearArticle:      clearArticle,
		ClearTmpFile:      clearTmpFile,
		ClearExpireServer: clearExpireServer,
	}
	cronCommand := cron2.NewCrontabCommand(subcommands)
	queueSubcommands := &queue.Subcommands{}
	queueCommand := queue.NewQueueCommand(queueSubcommands)
	exampleHandle := other.NewExampleHandle()
	exampleCommand := other2.NewExampleCommand(exampleHandle)
	otherSubcommands := &other2.Subcommands{
		ExampleCommand: exampleCommand,
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

var providerSet = wire.NewSet(provider.NewMySQLClient, provider.NewRedisClient, provider.NewHttpClient, provider.NewEmailClient, provider.NewRequestClient, filesystem.NewFilesystem, cache.NewSid, dao.NewBaseDao, cron2.NewCrontabCommand, cron.NewClearTmpFile, cron.NewClearArticle, cron.NewClearWsCache, cron.NewClearExpireServer, wire.Struct(new(cron2.Subcommands), "*"), queue.NewQueueCommand, wire.Struct(new(queue.Subcommands), "*"), queue2.NewEmailHandle, other2.NewOtherCommand, other2.NewExampleCommand, wire.Struct(new(other2.Subcommands), "*"), other.NewExampleHandle, wire.Struct(new(command.Commands), "*"), wire.Struct(new(AppProvider), "*"))
