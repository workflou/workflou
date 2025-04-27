<?php

namespace App\Filament\Resources\CrmAddresses\Tables;

use App\Filament\Exports\CrmAddressExporter;
use App\Filament\Imports\CrmAddressImporter;
use Filament\Actions\Action;
use Filament\Actions\BulkActionGroup;
use Filament\Actions\DeleteAction;
use Filament\Actions\DeleteBulkAction;
use Filament\Actions\EditAction;
use Filament\Actions\ExportAction;
use Filament\Actions\ExportBulkAction;
use Filament\Actions\ImportAction;
use Filament\Actions\ViewAction;
use Filament\Tables\Columns\TextColumn;
use Filament\Tables\Table;

class CrmAddressesTable
{
    public static function configure(Table $table): Table
    {
        return $table
            ->columns([
                TextColumn::make('address')
                    ->searchable(),
                TextColumn::make('city')
                    ->searchable(),
                TextColumn::make('state')
                    ->searchable(),
                TextColumn::make('zip')
                    ->searchable(),
                TextColumn::make('country')
                    ->searchable(),
            ])
            ->filters([
                //
            ])
            ->headerActions([
                ImportAction::make()->importer(CrmAddressImporter::class),
                ExportAction::make()->exporter(CrmAddressExporter::class),
            ])
            ->actions([
                EditAction::make(),
                DeleteAction::make(),
            ])
            ->bulkActions([
                BulkActionGroup::make([
                    DeleteBulkAction::make(),
                    ExportBulkAction::make()->exporter(CrmAddressExporter::class),
                ]),
            ]);
    }
}
