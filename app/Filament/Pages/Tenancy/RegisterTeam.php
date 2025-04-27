<?php

namespace App\Filament\Pages\Tenancy;

use App\Models\Team;
use App\TeamRole;
use App\TeamType;
use Filament\Actions\Action;
use Filament\Forms\Components\TextInput;
use Filament\Forms\Form;
use Filament\Pages\Tenancy\RegisterTenant;
use Filament\Schemas\Schema;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Str;

class RegisterTeam extends RegisterTenant
{
    public static function getLabel(): string
    {
        return __('New team');
    }

    public function form(Schema $schema): Schema
    {
        return $schema
            ->schema([
                TextInput::make('name')
                    ->required()
                    ->maxLength(255)
                    ->autofocus()
                    ->label(__('Team name')),
            ]);
    }

    public static function getSlug(): string
    {
        return 'teams/new';
    }

    public function getRegisterFormAction(): Action
    {
        return parent::getRegisterFormAction()
            ->label(__('Create'));
    }


    protected function handleRegistration(array $data): Model
    {
        $team = DB::transaction(function () use ($data) {
            $user = Auth::user();

            $team = $user->teams()->create([
                'name' => $name = $data['name'],
                'slug' => $slug = str(str($name)->slug() . '_' . Str::random(8))->lower(),
                'database' => $slug,
                'type' => TeamType::Company,
            ]);

            $team->users()->attach(Auth::user(), [
                'role' => TeamRole::OWNER,
            ]);

            $user->current_team_id = $team->id;
            $user->save();

            return $team;
        });

        return $team;
    }
}
