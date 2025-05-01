<?php

namespace App\Filament\Resources\CrmClients\Resources\CrmAddresses;

use App\Filament\Resources\CrmClients\CrmClientResource;
use App\Filament\Resources\CrmClients\Resources\CrmAddresses\Pages\CreateCrmAddress;
use App\Filament\Resources\CrmClients\Resources\CrmAddresses\Pages\EditCrmAddress;
use App\Filament\Resources\CrmClients\Resources\CrmAddresses\Pages\ViewCrmAddress;
use App\Filament\Resources\CrmClients\Resources\CrmAddresses\Schemas\CrmAddressForm;
use App\Filament\Resources\CrmClients\Resources\CrmAddresses\Schemas\CrmAddressInfolist;
use App\Filament\Resources\CrmClients\Resources\CrmAddresses\Tables\CrmAddressesTable;
use App\Models\CrmAddress;
use BackedEnum;
use Filament\Resources\Resource;
use Filament\Schemas\Schema;
use Filament\Support\Icons\Heroicon;
use Filament\Tables\Table;
use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\SoftDeletingScope;

class CrmAddressResource extends Resource
{
    protected static ?string $model = CrmAddress::class;

    protected static string|BackedEnum|null $navigationIcon = Heroicon::OutlinedRectangleStack;

    protected static ?string $parentResource = CrmClientResource::class;

    public static function form(Schema $schema): Schema
    {
        return CrmAddressForm::configure($schema);
    }

    public static function infolist(Schema $schema): Schema
    {
        return CrmAddressInfolist::configure($schema);
    }

    public static function table(Table $table): Table
    {
        return CrmAddressesTable::configure($table);
    }

    public static function getRelations(): array
    {
        return [
            //
        ];
    }

    public static function getPages(): array
    {
        return [
            'create' => CreateCrmAddress::route('/create'),
            'view' => ViewCrmAddress::route('/{record}'),
            'edit' => EditCrmAddress::route('/{record}/edit'),
        ];
    }

    public static function getEloquentQuery(): Builder
    {
        return parent::getEloquentQuery()
            ->withoutGlobalScopes([
                SoftDeletingScope::class,
            ]);
    }
}
