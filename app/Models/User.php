<?php

namespace App\Models;

use Illuminate\Contracts\Auth\MustVerifyEmail;

use App\Observers\UserObserver;
use Filament\Auth\MultiFactor\EmailCode\Contracts\HasEmailCodeAuthentication;
use Filament\Models\Contracts\HasDefaultTenant;
use Filament\Models\Contracts\HasTenants;
use Filament\Panel;
use Illuminate\Database\Eloquent\Attributes\ObservedBy;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Notifications\Notifiable;
use Illuminate\Support\Collection;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Str;

#[ObservedBy(UserObserver::class)]
class User extends Authenticatable implements HasTenants, HasDefaultTenant, HasEmailCodeAuthentication
{
    /** @use HasFactory<\Database\Factories\UserFactory> */
    use HasFactory, Notifiable;

    protected $fillable = [
        'name',
        'email',
        'password',
    ];

    protected $hidden = [
        'password',
        'remember_token',
    ];

    protected function casts(): array
    {
        return [
            'email_verified_at' => 'datetime',
            'password' => 'hashed',
        ];
    }

    public function teams(): HasMany
    {
        return $this->hasMany(Team::class);
    }

    public function currentTeam(): BelongsTo
    {
        return $this->belongsTo(Team::class, 'current_team_id');
    }

    public function switchTeam(Team $team): void
    {
        $this->current_team_id = $team->id;
        $this->save();
    }

    public function getTenants(Panel $panel): Collection
    {
        return $this->teams;
    }

    public function canAccessTenant(Model $tenant): bool
    {
        return $this->teams()->whereKey($tenant)->exists();
    }

    public function getDefaultTenant(Panel $panel): ?Model
    {
        return $this->currentTeam;
    }


    public function hasEmailCodeAuthentication(): bool
    {
        return false;
    }

    public function getEmailCodeAuthenticationSecret(): ?string
    {
        // return base32 string, hardcoded
        return 'JBSWY3DPEHPK3PXP';
    }

    public function saveEmailCodeAuthenticationSecret(?string $secret): void {}
}
