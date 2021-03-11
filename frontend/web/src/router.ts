import { createWebHistory, createRouter } from "vue-router";

const routes = [
	{
		path: "/",
		name: "Home",
		component: () => import("./views/home.vue"),
	},
	{
		path: "/login",
		name: "Login",
		component: () => import("./views/login.vue"),
	},
	{
		path: "/:catchAll(.*)",
		component: () => import("./views/not-found.vue"),
	},
];

const router = createRouter({
	routes,
	history: createWebHistory(),
});

router.beforeEach((to, from, next) => {
	const publicPages = ["/login", "/register"];
	const authRequired = !publicPages.includes(to.path);
	const loggedIn = localStorage.getItem("user");

	if (authRequired && !loggedIn) {
		return next("/login");
	}

	next();
});

export default router;
