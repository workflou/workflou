<?php

namespace App\Filament\Resources\CrmClients\Resources\CrmAddresses\Schemas;

use Filament\Infolists\Components\TextEntry;
use Filament\Schemas\Schema;

class CrmAddressInfolist
{
    public static function configure(Schema $schema): Schema
    {
        return $schema
            ->components([
                TextEntry::make('team.name')
                    ->numeric(),
                TextEntry::make('address'),
                TextEntry::make('city'),
                TextEntry::make('state'),
                TextEntry::make('zip'),
                TextEntry::make('country'),
                TextEntry::make('deleted_at')
                    ->dateTime(),
                TextEntry::make('created_at')
                    ->dateTime(),
                TextEntry::make('updated_at')
                    ->dateTime(),
            ]);
    }
}
