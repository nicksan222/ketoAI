package recipes_discover

/*
*
* GenerateDiscoverRecipes
* @param userId string
*
* This function generates the discover recipes for a user.
* The discover is cached in redis for 24 hours.
* The user has a limit of 100 recipes per day.
*
* The discover recipes are generated according to:
* 1. The user's preferences
* 2. The user's history
* 3. The user's following users
* 4. The user's following tags
*
 */
func GenerateDiscoverRecipes(
	userId string,
) {

}
