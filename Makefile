test:
	php artisan test

migrate-landlord:
	php artisan migrate --path=database/migrations/landlord --database=landlord

migrate-landlord-fresh:
	php artisan migrate:fresh --path=database/migrations/landlord --database=landlord
