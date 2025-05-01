<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;
use Illuminate\Database\Eloquent\SoftDeletes;

class CrmAddress extends Model
{
    /** @use HasFactory<\Database\Factories\CrmAddressFactory> */
    use HasFactory;
    use SoftDeletes;

    protected $fillable = [
        'address',
        'city',
        'state',
        'zip',
        'country',
    ];

    public function client(): BelongsTo
    {
        return $this->belongsTo(CrmClient::class);
    }

    public function leads(): BelongsToMany
    {
        return $this->belongsToMany(CrmLead::class);
    }

    public function proposals(): BelongsToMany
    {
        return $this->belongsToMany(CrmProposal::class);
    }

    public function team(): BelongsTo
    {
        return $this->belongsTo(Team::class);
    }
}
