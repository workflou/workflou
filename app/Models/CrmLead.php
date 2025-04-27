<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;

class CrmLead extends Model
{
    /** @use HasFactory<\Database\Factories\CrmLeadFactory> */
    use HasFactory;

    protected $fillable = [
        'name',
        'email',
        'phone',
        'status',
    ];

    public function clients(): BelongsToMany
    {
        return $this->belongsToMany(CrmClient::class);
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
