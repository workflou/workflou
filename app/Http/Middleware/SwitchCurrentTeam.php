<?php

namespace App\Http\Middleware;

use Closure;
use Filament\Facades\Filament;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Symfony\Component\HttpFoundation\Response;

class SwitchCurrentTeam
{
    public function handle(Request $request, Closure $next): Response
    {
        $team = Filament::getTenant();

        if ($team->id != Auth::user()->current_team_id && Auth::user()->can('switch', $team)) {
            Auth::user()->switchTeam($team);
        }

        return $next($request);
    }
}
