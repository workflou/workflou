<?php

namespace App\Filament\Resources\CrmAddresses\Pages;

use App\Filament\Resources\CrmAddresses\CrmAddressResource;
use Filament\Actions\EditAction;
use Filament\Resources\Pages\ViewRecord;

class ViewCrmAddress extends ViewRecord
{
    protected static string $resource = CrmAddressResource::class;

    protected function getHeaderActions(): array
    {
        return [
            EditAction::make(),
        ];
    }
}
