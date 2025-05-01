<?php

namespace App\Filament\Resources\CrmClients\Pages;

use App\Filament\Resources\CrmClients\CrmClientResource;
use Filament\Actions\CreateAction;
use Filament\Resources\Pages\ListRecords;

class ListCrmClients extends ListRecords
{
    protected static string $resource = CrmClientResource::class;

    protected function getHeaderActions(): array
    {
        return [
            CreateAction::make(),
        ];
    }
}
