<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;
use Illuminate\Database\Eloquent\SoftDeletes;

class CrmProposal extends Model
{
    /** @use HasFactory<\Database\Factories\CrmProposalFactory> */
    use HasFactory;
    use SoftDeletes;

    protected $fillable = [
        'number',
        'title',
        'description',
        'status',
        'amount',
        'due_date',
    ];

    public function clients(): BelongsToMany
    {
        return $this->belongsToMany(CrmClient::class);
    }

    public function leads(): BelongsToMany
    {
        return $this->belongsToMany(CrmLead::class);
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
