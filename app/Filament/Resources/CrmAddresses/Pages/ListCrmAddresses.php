<?php

namespace App\Filament\Resources\CrmAddresses\Pages;

use App\Filament\Resources\CrmAddresses\CrmAddressResource;
use Filament\Actions\CreateAction;
use Filament\Resources\Pages\ListRecords;

class ListCrmAddresses extends ListRecords
{
    protected static string $resource = CrmAddressResource::class;

    protected function getHeaderActions(): array
    {
        return [
            CreateAction::make(),
        ];
    }
}
