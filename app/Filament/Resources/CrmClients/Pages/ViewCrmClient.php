<?php

namespace App\Filament\Resources\CrmClients\Pages;

use App\Filament\Resources\CrmClients\CrmClientResource;
use Filament\Actions\EditAction;
use Filament\Resources\Pages\ViewRecord;

class ViewCrmClient extends ViewRecord
{
    protected static string $resource = CrmClientResource::class;

    protected function getHeaderActions(): array
    {
        return [
            EditAction::make(),
        ];
    }
}
