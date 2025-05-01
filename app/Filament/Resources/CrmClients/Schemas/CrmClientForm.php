<?php

namespace App\Filament\Resources\CrmClients\Schemas;

use Filament\Forms\Components\TextInput;
use Filament\Schemas\Schema;

class CrmClientForm
{
    public static function configure(Schema $schema): Schema
    {
        return $schema
            ->components([
                TextInput::make('name')
                    ->label(__('Name'))
                    ->required()
                    ->columnSpanFull()
                    ->maxLength(255)
                    ->placeholder(__('Enter name')),
                TextInput::make('email')
                    ->label(__('Email'))
                    ->email()
                    ->maxLength(255)
                    ->placeholder(__('Enter email')),
                TextInput::make('phone')
                    ->label(__('Phone'))
                    ->maxLength(255)
                    ->tel()
                    ->placeholder(__('Enter phone')),
            ]);
    }
}
