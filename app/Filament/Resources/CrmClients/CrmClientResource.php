<?php

namespace App\Filament\Resources\CrmClients;

use App\Filament\Resources\CrmClients\Pages\CreateCrmClient;
use App\Filament\Resources\CrmClients\Pages\EditCrmClient;
use App\Filament\Resources\CrmClients\Pages\ListCrmClients;
use App\Filament\Resources\CrmClients\Pages\ViewCrmClient;
use App\Filament\Resources\CrmClients\Schemas\CrmClientForm;
use App\Filament\Resources\CrmClients\Schemas\CrmClientInfolist;
use App\Filament\Resources\CrmClients\Tables\CrmClientsTable;
use App\Models\CrmClient as ModelsCrmClient;
use BackedEnum;
use CrmClient;
use Filament\Resources\Resource;
use Filament\Schemas\Schema;
use Filament\Support\Icons\Heroicon;
use Filament\Tables\Table;
use Illuminate\Contracts\Support\Htmlable;
use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletingScope;
use UnitEnum;

class CrmClientResource extends Resource
{
    protected static ?string $model = ModelsCrmClient::class;

    protected static string|BackedEnum|null $navigationIcon = Heroicon::OutlinedRectangleStack;

    protected static ?string $label = 'Client';

    protected static ?string $pluralLabel = 'Clients';

    protected static ?string $slug = 'clients';

    protected static string | UnitEnum | null $navigationGroup = 'CRM';

    protected static bool $isGloballySearchable = true;

    public static function getGloballySearchableAttributes(): array
    {
        return ['name', 'email', 'phone'];
    }

    public static function getGlobalSearchResultTitle(Model $record): string | Htmlable
    {
        return $record->name;
    }

    public static function getGlobalSearchResultDetails(Model $record): array
    {
        return [
            __('Email') => $record->email,
            __('Phone') => $record->phone,
        ];
    }

    public static function form(Schema $schema): Schema
    {
        return CrmClientForm::configure($schema);
    }

    public static function infolist(Schema $schema): Schema
    {
        return CrmClientInfolist::configure($schema);
    }

    public static function table(Table $table): Table
    {
        return CrmClientsTable::configure($table);
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
            'index' => ListCrmClients::route('/'),
            'edit' => EditCrmClient::route('/{record}/edit'),
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
