package com.lomolo.copodapp.ui.navigation

import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.navigation.NavGraph.Companion.findStartDestination
import androidx.navigation.NavHostController
import androidx.navigation.NavType
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.navArgument
import com.lomolo.copodapp.ui.screens.ExploreMarketsScreen
import com.lomolo.copodapp.ui.screens.ExploreMarketsScreenDestination
import com.lomolo.copodapp.ui.screens.HomeScreen
import com.lomolo.copodapp.ui.screens.HomeScreenDestination
import com.lomolo.copodapp.ui.screens.LandScreen
import com.lomolo.copodapp.ui.screens.LandScreenDestination
import com.lomolo.copodapp.ui.screens.LoadingScreen
import com.lomolo.copodapp.ui.screens.LoadingScreenDestination
import com.lomolo.copodapp.ui.screens.LoginScreen
import com.lomolo.copodapp.ui.screens.LoginScreenDestination
import com.lomolo.copodapp.ui.screens.MpesaScreen
import com.lomolo.copodapp.ui.screens.MpesaScreenDestination
import com.lomolo.copodapp.ui.screens.UploadGovtIssuedId
import com.lomolo.copodapp.ui.screens.UploadGovtIssuedIdScreenDestination
import com.lomolo.copodapp.ui.screens.UploadLandTitle
import com.lomolo.copodapp.ui.screens.UploadLandTitleScreenDestination
import com.lomolo.copodapp.ui.screens.Web3SdkErrorScreen
import com.lomolo.copodapp.ui.screens.Web3SdkErrorScreenDestination
import com.lomolo.copodapp.ui.viewmodels.InitializeSdk
import com.lomolo.copodapp.ui.viewmodels.MainViewModel
import com.lomolo.copodapp.ui.viewmodels.MpesaViewModel
import com.lomolo.copodapp.ui.viewmodels.RegisterLandViewModel
import org.koin.androidx.compose.navigation.koinNavViewModel

interface Navigation {
    val title: Int?
    val route: String
}

object RootNavigation : Navigation {
    override val title = null
    override val route = "root"
}

@Composable
fun NavigationHost(
    modifier: Modifier,
    navHostController: NavHostController,
    mainViewModel: MainViewModel,
) {
    val isLoggedIn by mainViewModel.isLoggedIn.collectAsState()
    val startRoute = when (mainViewModel.initializeSdk) {
        InitializeSdk.Loading -> {
            LoadingScreenDestination.route
        }

        InitializeSdk.Success -> {
            when (isLoggedIn) {
                true -> ExploreMarketsScreenDestination.route
                false -> HomeScreenDestination.route
            }
        }

        else -> Web3SdkErrorScreenDestination.route
    }
    val onNavigateTo = { route: String ->
        navHostController.navigate(route) {
            // Pop up to the start destination of the graph to
            // avoid building up a large stack of destinations
            // on the back stack as users select items
            popUpTo(navHostController.graph.findStartDestination().id) {
                saveState = false
            }
            // Avoid multiple copies of the same destination when
            // re-selecting the same item
            launchSingleTop = true
            // Restore state when re-selecting a previously selected item
            restoreState = true
        }
    }

    NavHost(
        navController = navHostController,
        startDestination = startRoute.toString(),
        route = RootNavigation.route,
    ) {
        composable(route = LoginScreenDestination.route) {
            LoginScreen(mainViewModel = mainViewModel, onGoBack = {
                navHostController.popBackStack()
            })
        }
        composable(route = ExploreMarketsScreenDestination.route) {
            ExploreMarketsScreen(
                mainViewModel = mainViewModel,
                onNavigateTo = onNavigateTo,
                currentDestination = it.destination,
            )
        }
        composable(route = LoadingScreenDestination.route) {
            LoadingScreen(
                modifier = modifier,
            )
        }
        composable(route = Web3SdkErrorScreenDestination.route) {
            Web3SdkErrorScreen(
                modifier = modifier,
            )
        }
        composable(route = HomeScreenDestination.route) {
            HomeScreen(modifier = modifier, onGoToLogin = {
                navHostController.navigate(LoginScreenDestination.route)
            })
        }
        composable(route = LandScreenDestination.route) {
            LandScreen(
                onNavigateTo = onNavigateTo,
                currentDestination = it.destination,
                userInfo = mainViewModel.userInfo,
                mainViewModel = mainViewModel,
                onClickAddLand = {
                    navHostController.navigate(UploadLandTitleScreenDestination.route)
                }
            )
        }
        composable(route = UploadLandTitleScreenDestination.route) {
            val registerLandViewModel: RegisterLandViewModel = koinNavViewModel()
            UploadLandTitle(
                onGoBack = {
                    navHostController.popBackStack()
                },
                onNavigateTo = {
                    navHostController.navigate(it)
                },
                userEmail = mainViewModel.userInfo!!.email,
                userWallet = mainViewModel.credentials!!.address,
                viewModel = registerLandViewModel,
            )
        }
        composable(
            route = UploadGovtIssuedIdScreenDestination.routeWithArgs,
            arguments = listOf(navArgument(UploadGovtIssuedIdScreenDestination.LAND_TITLE_ID_ARG) {
                type = NavType.StringType
            })
        ) {
            val registerLandViewModel: RegisterLandViewModel = koinNavViewModel()
            UploadGovtIssuedId(
                onGoBack = {
                    navHostController.popBackStack()
                },
                userEmail = mainViewModel.userInfo!!.email,
                userWallet = mainViewModel.credentials!!.address,
                viewModel = registerLandViewModel,
                onNext = { uploadId ->
                    navHostController.navigate("${MpesaScreenDestination.route}/${uploadId}")
                },
            )
        }
        composable(
            route = MpesaScreenDestination.routeWithArgs,
            arguments = listOf(navArgument(MpesaScreenDestination.LAND_TITLE_ID_ARG) {
                type = NavType.StringType
            })
        ) {
            val mpesaViewModel: MpesaViewModel = koinNavViewModel()
            MpesaScreen(
                onGoBack = {
                    navHostController.popBackStack()
                },
                viewModel = mpesaViewModel,
                mainViewModel = mainViewModel,
            )
        }
    }
}