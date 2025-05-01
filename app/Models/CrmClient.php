<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;
use Illuminate\Database\Eloquent\SoftDeletes;

class CrmClient extends Model
{
    /** @use HasFactory<\Database\Factories\CrmClientFactory> */
    use HasFactory;
    use SoftDeletes;

    protected $fillable = [
        'name',
        'email',
        'phone',
    ];

    public function leads(): BelongsToMany
    {
        return $this->belongsToMany(CrmLead::class);
    }

    public function proposals(): BelongsToMany
    {
        return $this->belongsToMany(CrmProposal::class);
    }

    public function addresses(): BelongsToMany
    {
        return $this->belongsToMany(CrmAddress::class);
    }

    public function users(): BelongsToMany
    {
        return $this->belongsToMany(User::class);
    }

    public function team(): BelongsTo
    {
        return $this->belongsTo(Team::class);
    }
}
