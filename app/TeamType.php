<?php

namespace App;

use Filament\Support\Contracts\HasLabel;

enum TeamType: string implements HasLabel
{
    case Personal = 'personal';
    case Company = 'company';

    public function getLabel(): ?string
    {
        return match ($this) {
            self::Personal => __('Personal'),
            self::Company => __('Company'),
        };
    }
}
