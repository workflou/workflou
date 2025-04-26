<?php

use App\Models\Team;
use App\Models\User;
use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    public function up(): void
    {
        Schema::create('teams', function (Blueprint $table) {
            $table->id();
            $table->string('name');
            $table->string('slug')->unique();
            $table->string('database')->unique();
            $table->string('type')->default('personal');
            $table->foreignIdFor(User::class)->constrained();
            $table->timestamps();
        });

        Schema::create('team_user', function (Blueprint $table) {
            $table->id();
            $table->string('role')->default('member');
            $table->foreignIdFor(Team::class)->constrained();
            $table->foreignIdFor(User::class)->constrained();
            $table->timestamps();
        });

        Schema::table('users', function (Blueprint $table) {
            $table->foreignIdFor(Team::class, 'current_team_id')->nullable()->constrained()->nullOnDelete();
        });
    }

    public function down(): void
    {
        Schema::dropIfExists('teams');
        Schema::dropIfExists('team_user');
        Schema::dropIfExists('users');
    }
};
