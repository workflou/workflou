<?php

namespace App\Filament\Resources\CrmClients\Schemas;

use Filament\Infolists\Components\TextEntry;
use Filament\Schemas\Schema;

class CrmClientInfolist
{
    public static function configure(Schema $schema): Schema
    {
        return $schema
            ->components([
                TextEntry::make('name')
                    ->label(__('Name'))
                    ->columnSpan(2),
                TextEntry::make('email')
                    ->label(__('Email'))
                    ->columnSpan(1),
                TextEntry::make('phone')
                    ->label(__('Phone'))
                    ->columnSpan(1),
            ]);
    }
}
