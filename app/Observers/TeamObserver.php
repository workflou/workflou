<?php

namespace App\Observers;

use App\Models\Team;
use Illuminate\Support\Str;

class TeamObserver
{
    public function creating(Team $team): void
    {
        $team->slug = str(str($team->name)->slug() . '-' . Str::random(8))->lower();
        $team->database = $team->slug;
    }
}
