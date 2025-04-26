<?php

namespace App\Models;

use App\Observers\TeamObserver;
use App\TeamType;
use Illuminate\Database\Eloquent\Attributes\ObservedBy;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Spatie\Multitenancy\Contracts\IsTenant;
use Spatie\Multitenancy\Models\Concerns\ImplementsTenant;
use Spatie\Multitenancy\Models\Concerns\UsesLandlordConnection;

#[ObservedBy(TeamObserver::class)]
class Team extends Model implements IsTenant
{
    /** @use HasFactory<\Database\Factories\TeamFactory> */
    use HasFactory;
    use UsesLandlordConnection;
    use ImplementsTenant;

    protected $fillable = [
        'name',
        'slug',
        'type',
        'database',
    ];

    protected $casts = [
        'type' => TeamType::class,
    ];

    public function getRouteKeyName(): string
    {
        return 'slug';
    }
}
