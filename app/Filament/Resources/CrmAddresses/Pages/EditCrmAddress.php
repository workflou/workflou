<?php

namespace App\Filament\Resources\CrmAddresses\Pages;

use App\Filament\Resources\CrmAddresses\CrmAddressResource;
use Filament\Actions\DeleteAction;
use Filament\Actions\ViewAction;
use Filament\Resources\Pages\EditRecord;

class EditCrmAddress extends EditRecord
{
    protected static string $resource = CrmAddressResource::class;

    protected function getHeaderActions(): array
    {
        return [
            ViewAction::make(),
            DeleteAction::make(),
        ];
    }
}
