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
            'slug' => $slug = str(str($user->name)->slug() . '_' . Str::random(8))->lower(),
            'type' => TeamType::Personal,
            'database' => $slug,
        ]);

        $user->switchTeam($team);
    }
}
