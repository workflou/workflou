<?php

namespace App\Providers\Filament;

use App\Filament\Pages\Tenancy\RegisterTeam;
use App\Http\Middleware\SwitchCurrentTeam;
use App\Models\Team;
use Filament\Actions\Action;
use Filament\Auth\MultiFactor\EmailCode\EmailCodeAuthentication;
use Filament\Http\Middleware\Authenticate;
use Filament\Http\Middleware\AuthenticateSession;
use Filament\Http\Middleware\DisableBladeIconComponents;
use Filament\Http\Middleware\DispatchServingFilamentEvent;
use Filament\Navigation\MenuItem;
use Filament\Pages\Dashboard;
use Filament\Panel;
use Filament\PanelProvider;
use Filament\Support\Colors\Color;
use Filament\Support\Enums\Width;
use Filament\Widgets\AccountWidget;
use Filament\Widgets\FilamentInfoWidget;
use Illuminate\Cookie\Middleware\AddQueuedCookiesToResponse;
use Illuminate\Cookie\Middleware\EncryptCookies;
use Illuminate\Foundation\Http\Middleware\VerifyCsrfToken;
use Illuminate\Routing\Middleware\SubstituteBindings;
use Illuminate\Session\Middleware\StartSession;
use Illuminate\View\Middleware\ShareErrorsFromSession;

class AppPanelProvider extends PanelProvider
{
    public function panel(Panel $panel): Panel
    {
        return $panel
            ->default()
            ->id('app')
            ->path('')
            ->tenant(Team::class, slugAttribute: 'slug')
            ->tenantRoutePrefix('teams')
            ->tenantMiddleware([
                SwitchCurrentTeam::class,
            ])
            ->spa()
            ->topNavigation()
            ->brandLogo(fn() => view('filament.app.logo'))
            ->viteTheme('resources/css/filament/app/theme.css')
            ->maxContentWidth(Width::Full)
            ->login()
            // ->emailVerification()
            // ->requiresMultiFactorAuthentication()
            // ->multiFactorAuthentication([
            //     EmailCodeAuthentication::make(),
            // ])
            ->registration()
            ->tenantRegistration(RegisterTeam::class)
            ->colors([
                'primary' => Color::Violet,
            ])
            ->discoverResources(in: app_path('Filament/Resources'), for: 'App\Filament\Resources')
            ->discoverClusters(in: app_path('Filament/Clusters'), for: 'App\Filament\Clusters')
            ->discoverPages(in: app_path('Filament/Pages'), for: 'App\Filament\Pages')
            ->pages([
                Dashboard::class,
            ])
            ->discoverWidgets(in: app_path('Filament/Widgets'), for: 'App\Filament\Widgets')
            ->widgets([
                AccountWidget::class,
                FilamentInfoWidget::class,
            ])
            ->middleware([
                EncryptCookies::class,
                AddQueuedCookiesToResponse::class,
                StartSession::class,
                AuthenticateSession::class,
                ShareErrorsFromSession::class,
                VerifyCsrfToken::class,
                SubstituteBindings::class,
                DisableBladeIconComponents::class,
                DispatchServingFilamentEvent::class,
            ])
            ->authMiddleware([
                Authenticate::class,
            ]);
    }
}
