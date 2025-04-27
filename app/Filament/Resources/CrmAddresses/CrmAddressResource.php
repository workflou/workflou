<?php

namespace App\Filament\Resources\CrmAddresses;

use App\Filament\Clusters\Crm\CrmCluster;
use App\Filament\Resources\CrmAddresses\Pages\CreateCrmAddress;
use App\Filament\Resources\CrmAddresses\Pages\EditCrmAddress;
use App\Filament\Resources\CrmAddresses\Pages\ListCrmAddresses;
use App\Filament\Resources\CrmAddresses\Pages\ViewCrmAddress;
use App\Filament\Resources\CrmAddresses\Schemas\CrmAddressForm;
use App\Filament\Resources\CrmAddresses\Schemas\CrmAddressInfolist;
use App\Filament\Resources\CrmAddresses\Tables\CrmAddressesTable;
use App\Models\CrmAddress;
use BackedEnum;
use Closure;
use Filament\Resources\Resource;
use Filament\Schemas\Schema;
use Filament\Support\Icons\Heroicon;
use Filament\Tables\Table;
use Illuminate\Contracts\Support\Htmlable;
use Illuminate\Database\Eloquent\Model;
use UnitEnum;

class CrmAddressResource extends Resource
{
    protected static ?string $model = CrmAddress::class;

    protected static string|BackedEnum|null $navigationIcon = Heroicon::OutlinedRectangleStack;

    protected static ?string $label = 'Address';

    protected static ?string $pluralLabel = 'Addresses';

    protected static ?string $slug = 'addresses';

    protected static string | UnitEnum | null $navigationGroup = 'CRM';

    protected static bool $isGloballySearchable = true;

    public static function getGloballySearchableAttributes(): array
    {
        return ['address', 'city', 'state', 'zip', 'country'];
    }

    public static function getGlobalSearchResultTitle(Model $record): string | Htmlable
    {
        return $record->address . ' ' . $record->city . ' ' . $record->state . ' ' . $record->zip . ' ' . $record->country;
    }

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
            'index' => ListCrmAddresses::route('/'),
        ];
    }
}
