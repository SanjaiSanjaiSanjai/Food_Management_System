package baseurl

var REGISTER_URL = map[string]string{"register": "/register"}

var LOGIN_URL = map[string]string{"login": "/login"}

var USER_ADDRESS_URL = map[string]string{"user_address": "/user/address"}

var RESTAURANT_URL = map[string]string{"restaurant": "/restaurant/create", "getRestaurant": "/restaurant/getRestaurants", "restaurantAddress": "/restaurant/create/address"}

var MENU_URL = map[string]string{"menu": "/menu/create/:restaurant_id", "menu_category": "/menu/category/create/:restaurant_id", "getMenuCategories": "/menu/category/get", "getCategoryById": "/menu/category/get/:category_id"}

var RESTAURANT_CATEGORY_URL = map[string]string{"restaurant_category": "/restaurant/category/create/:restaurant_id"}
