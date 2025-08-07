# Variables
COUNT = 10

# Seed random products
seed-random:
	go run cmd/seeder/main.go -mode random -count $(COUNT)

# Seed men's clothing
seed-mens:
	go run cmd/seeder/main.go -mode category -category "Men's Clothing" -count $(COUNT)

# Seed diverse mix
seed-diverse:
	go run cmd/seeder/main.go -mode diverse -count $(COUNT)

# Seed batch
seed-batch:
	go run cmd/seeder/main.go -mode batch -count $(COUNT)
