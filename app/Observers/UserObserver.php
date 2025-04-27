<?php

namespace App\Observers;

use App\Models\User;
use App\TeamType;
use Illuminate\Support\Str;

class UserObserver
{
    public function created(User $user): void
    {
        $team = $user->teams()->create([
            'name' => $user->name . '\'s Team',
            'type' => TeamType::Personal,
        ]);

        $user->switchTeam($team);
    }
}
