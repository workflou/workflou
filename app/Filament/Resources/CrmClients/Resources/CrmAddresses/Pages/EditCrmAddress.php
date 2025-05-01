<?php

namespace App\Filament\Resources\CrmClients\Resources\CrmAddresses\Pages;

use App\Filament\Resources\CrmClients\Resources\CrmAddresses\CrmAddressResource;
use Filament\Actions\DeleteAction;
use Filament\Actions\ForceDeleteAction;
use Filament\Actions\RestoreAction;
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
            ForceDeleteAction::make(),
            RestoreAction::make(),
        ];
    }
}
