<?php

use App\Models\User;
use Filament\Auth\Pages\Register;
use Livewire\Livewire;

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
