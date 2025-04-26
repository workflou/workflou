<?php

namespace App;

use Filament\Support\Contracts\HasLabel;

enum TeamRole: string implements HasLabel
{
    case OWNER = 'owner';
    case MEMBER = 'member';

    public function getLabel(): ?string
    {
        return match ($this) {
            self::OWNER => __('Owner'),
            self::MEMBER => __('Member'),
        };
    }
}
