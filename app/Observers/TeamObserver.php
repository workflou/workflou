<?php

namespace App\Observers;

use App\Models\Team;

class TeamObserver
{
    public function created(Team $team): void
    {
        touch(database_path($team->database));
    }
}
