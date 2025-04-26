<?php

use App\Models\Team;
use App\Models\User;

test('user has a default team', function () {
    $user = User::factory()->create();

    expect($user->currentTeam)->not->toBeNull();
});

test('user can switch team', function () {
    $user = User::factory()->create();
    $team = Team::factory()->create();
    $user->switchTeam($team);

    expect($user->fresh()->currentTeam->id)->toBe($team->id);
});
