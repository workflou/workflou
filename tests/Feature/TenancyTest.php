<?php

use App\Filament\Pages\Tenancy\RegisterTeam;
use App\Models\Team;
use App\Models\User;
use Filament\Auth\Pages\Register;
use Livewire\Livewire;

use function Pest\Laravel\actingAs;
use function Pest\Laravel\assertDatabaseHas;

test('new user has a default team', function () {
    Livewire::test(Register::class)
        ->fillForm([
            'name' => 'John Doe',
            'email' => 'john@doe.com',
            'password' => 'password',
            'passwordConfirmation' => 'password',
        ])
        ->call('register')
        ->assertHasNoErrors();
    expect(User::first()->currentTeam)->not->toBeNull();
});

test('user can create a new team', function () {
    actingAs($user = User::factory()->create());

    Livewire::test(RegisterTeam::class)
        ->assertSuccessful()
        ->fillForm([
            'name' => 'John Doe',
        ])
        ->call('register')
        ->assertHasNoErrors();

    expect($user->teams()->count())->toBe(2);

    assertDatabaseHas('teams', [
        'name' => 'John Doe',
        'user_id' => $user->id,
    ]);


    expect(Team::find(2)->users()->count())->toBe(1);
});
