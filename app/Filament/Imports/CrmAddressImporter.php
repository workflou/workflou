<?php

namespace App\Filament\Imports;

use App\Models\CrmAddress;
use Filament\Actions\Imports\ImportColumn;
use Filament\Actions\Imports\Importer;
use Filament\Actions\Imports\Models\Import;
use Illuminate\Support\Number;

class CrmAddressImporter extends Importer
{
    protected static ?string $model = CrmAddress::class;

    public static function getColumns(): array
    {
        return [
            ImportColumn::make('address'),
            ImportColumn::make('city'),
            ImportColumn::make('state'),
            ImportColumn::make('zip'),
            ImportColumn::make('country'),
        ];
    }

    public function resolveRecord(): CrmAddress
    {
        return new CrmAddress();
    }

    public static function getCompletedNotificationBody(Import $import): string
    {
        $body = 'Your crm address import has completed and ' . Number::format($import->successful_rows) . ' ' . str('row')->plural($import->successful_rows) . ' imported.';

        if ($failedRowsCount = $import->getFailedRowsCount()) {
            $body .= ' ' . Number::format($failedRowsCount) . ' ' . str('row')->plural($failedRowsCount) . ' failed to import.';
        }

        return $body;
    }
}
