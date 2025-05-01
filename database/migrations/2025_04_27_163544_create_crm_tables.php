<?php

use App\Models\CrmAddress;
use App\Models\CrmClient;
use App\Models\CrmLead;
use App\Models\CrmProposal;
use App\Models\Team;
use App\Models\User;
use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    public function up(): void
    {
        Schema::create('crm_addresses', function (Blueprint $table) {
            $table->id();
            $table->foreignIdFor(Team::class)->constrained();
            $table->string('address')->nullable();
            $table->string('city')->nullable();
            $table->string('state')->nullable();
            $table->string('zip')->nullable();
            $table->string('country')->nullable();
            $table->softDeletes();
            $table->timestamps();
        });

        Schema::create('crm_clients', function (Blueprint $table) {
            $table->id();
            $table->foreignIdFor(Team::class)->constrained();
            $table->string('name');
            $table->string('email')->nullable();
            $table->string('phone')->nullable();
            $table->softDeletes();
            $table->timestamps();
        });

        Schema::create('crm_leads', function (Blueprint $table) {
            $table->id();
            $table->foreignIdFor(Team::class)->constrained();
            $table->string('name');
            $table->string('email');
            $table->string('phone');
            $table->string('status');
            $table->softDeletes();
            $table->timestamps();
        });

        Schema::create('crm_proposals', function (Blueprint $table) {
            $table->id();
            $table->foreignIdFor(Team::class)->constrained();
            $table->string('number');
            $table->foreignIdFor(CrmClient::class)->constrained();
            $table->foreignIdFor(CrmLead::class)->constrained();
            $table->foreignIdFor(User::class)->constrained();
            $table->text('title')->nullable();
            $table->longText('description')->nullable();
            $table->string('status')->nullable();
            $table->bigInteger('amount')->nullable();
            $table->timestamp('due_date')->nullable();
            $table->softDeletes();
            $table->timestamps();
        });

        Schema::create('crm_address_client', function (Blueprint $table) {
            $table->id();
            $table->foreignIdFor(CrmAddress::class)->constrained();
            $table->foreignIdFor(CrmClient::class)->constrained();
            $table->timestamps();
        });

        Schema::create('crm_lead_client', function (Blueprint $table) {
            $table->id();
            $table->foreignIdFor(CrmLead::class)->constrained();
            $table->foreignIdFor(CrmClient::class)->constrained();
            $table->timestamps();
        });

        Schema::create('crm_lead_user', function (Blueprint $table) {
            $table->id();
            $table->foreignIdFor(CrmLead::class)->constrained();
            $table->foreignIdFor(User::class)->constrained();
            $table->timestamps();
        });

        Schema::create('crm_proposal_user', function (Blueprint $table) {
            $table->id();
            $table->foreignIdFor(CrmProposal::class)->constrained();
            $table->foreignIdFor(User::class)->constrained();
            $table->timestamps();
        });
    }

    public function down(): void
    {
        Schema::dropIfExists('crm_address_client');
        Schema::dropIfExists('crm_lead_client');
        Schema::dropIfExists('crm_leads');
        Schema::dropIfExists('crm_clients');
        Schema::dropIfExists('crm_addresses');
        Schema::dropIfExists('crm_proposal_user');
        Schema::dropIfExists('crm_proposals');
        Schema::dropIfExists('crm_lead_user');
    }
};
