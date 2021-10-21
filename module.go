package core

import "go.uber.org/fx"

var Module = fx.Options(
	LogrusModule,
	ClerkModule,
	MessengerModule,
	EchoModule,
	GraphQlModule,
	fx.Invoke(
		UseEnvironment,
	),
)
