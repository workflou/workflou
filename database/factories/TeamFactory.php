<?php

namespace Database\Factories;

use App\Models\User;
use App\TeamType;
use Illuminate\Database\Eloquent\Factories\Factory;

class TeamFactory extends Factory
{
    public function definition(): array
    {
        return [
            'name' => fake()->name(),
            'slug' => fake()->slug(),
            'type' => TeamType::Personal,
            'database' => fake()->word(),
            'user_id' => User::factory(),
        ];
    }
}
