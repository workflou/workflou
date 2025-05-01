<?php

namespace App\Filament\Resources\CrmClients\Resources\CrmAddresses\Schemas;

use Filament\Forms\Components\Select;
use Filament\Forms\Components\TextInput;
use Filament\Schemas\Schema;

class CrmAddressForm
{
    public static function configure(Schema $schema): Schema
    {
        return $schema
            ->components([
                Select::make('team_id')
                    ->relationship('team', 'name')
                    ->required(),
                TextInput::make('address'),
                TextInput::make('city'),
                TextInput::make('state'),
                TextInput::make('zip'),
                TextInput::make('country'),
            ]);
    }
}
