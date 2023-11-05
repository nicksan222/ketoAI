insert-ingredients:
	echo "Inserting ingredients..."
	bun run ./ingredients/insert.ts

generate-recipes:
	for i in $$(seq 1 $(TIMES)); do \
		echo "Generating recipe $$i..."; \
		bun run ./recipes/generate_recipe.ts; \
	done
