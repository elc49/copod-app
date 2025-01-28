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
import com.lomolo.copodapp.state.viewmodels.InitializeSdk
import com.lomolo.copodapp.state.viewmodels.LandTitleDetailsViewModel
import com.lomolo.copodapp.state.viewmodels.MainViewModel
import com.lomolo.copodapp.state.viewmodels.MpesaViewModel
import com.lomolo.copodapp.state.viewmodels.OnboardingViewModel
import com.lomolo.copodapp.ui.screens.AccountScreen
import com.lomolo.copodapp.ui.screens.AccountScreenDestination
import com.lomolo.copodapp.ui.screens.CreateLandScreen
import com.lomolo.copodapp.ui.screens.CreateLandScreenDestination
import com.lomolo.copodapp.ui.screens.ErrorScreen
import com.lomolo.copodapp.ui.screens.ErrorScreenDestination
import com.lomolo.copodapp.ui.screens.ExploreMarketsScreen
import com.lomolo.copodapp.ui.screens.ExploreMarketsScreenDestination
import com.lomolo.copodapp.ui.screens.FoundLandScreen
import com.lomolo.copodapp.ui.screens.FoundLandScreenDestination
import com.lomolo.copodapp.ui.screens.HomeScreen
import com.lomolo.copodapp.ui.screens.HomeScreenDestination
import com.lomolo.copodapp.ui.screens.LandScreen
import com.lomolo.copodapp.ui.screens.LandScreenDestination
import com.lomolo.copodapp.ui.screens.LoadingScreen
import com.lomolo.copodapp.ui.screens.LoadingScreenDestination
import com.lomolo.copodapp.ui.screens.MpesaScreen
import com.lomolo.copodapp.ui.screens.MpesaScreenDestination
import com.lomolo.copodapp.ui.screens.OnboardingScreen
import com.lomolo.copodapp.ui.screens.OnboardingScreenDestination
import com.lomolo.copodapp.ui.screens.SearchLandScreen
import com.lomolo.copodapp.ui.screens.SearchScreenDestination
import com.lomolo.copodapp.ui.screens.SuccessScreen
import com.lomolo.copodapp.ui.screens.SuccessScreenDestination
import com.lomolo.copodapp.ui.screens.UploadDisplayPicture
import com.lomolo.copodapp.ui.screens.UploadDisplayPictureDestination
import com.lomolo.copodapp.ui.screens.UploadGovtIssuedId
import com.lomolo.copodapp.ui.screens.UploadGovtIssuedIdScreenDestination
import com.lomolo.copodapp.ui.screens.UploadLandTitle
import com.lomolo.copodapp.ui.screens.UploadLandTitleScreenDestination
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
    onboardingViewModel: OnboardingViewModel,
    landTitleViewModel: LandTitleDetailsViewModel,
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

        else -> ErrorScreenDestination.route
    }
    val onNext = { route: String ->
        navHostController.navigate(route) { launchSingleTop = true }
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
        composable(route = ExploreMarketsScreenDestination.route) {
            ExploreMarketsScreen(
                onNavigateTo = onNavigateTo,
                mainViewModel = mainViewModel,
                onNext = {
                    onNext(it)
                },
                currentDestination = it.destination,
            )
        }
        composable(route = LoadingScreenDestination.route) {
            LoadingScreen(
                modifier = modifier,
            )
        }
        composable(route = ErrorScreenDestination.route) {
            ErrorScreen(
                modifier = modifier,
            )
        }
        composable(route = HomeScreenDestination.route) {
            HomeScreen(modifier = modifier, mainViewModel = mainViewModel)
        }
        composable(route = LandScreenDestination.route) {
            LandScreen(
                userInfo = mainViewModel.userInfo,
                onGoBack = {
                    navHostController.popBackStack()
                })
        }
        composable(
            route = UploadLandTitleScreenDestination.routeWithArgs,
            arguments = listOf(navArgument(UploadLandTitleScreenDestination.RESUBMIT_ID_ARG) {
                type = NavType.BoolType
            })
        ) {
            val isResubmit = it.arguments?.getBoolean("reSubmit", false)

            UploadLandTitle(
                isResubmit = isResubmit,
                onGoBack = {
                    navHostController.popBackStack()
                },
                onNavigateTo = {
                    onNext(it)
                },
                viewModel = onboardingViewModel,
            )
        }
        composable(
            route = UploadGovtIssuedIdScreenDestination.routeWithArgs,
            arguments = listOf(navArgument(UploadGovtIssuedIdScreenDestination.RESUBMIT_ID_ARG) {
                type = NavType.BoolType
            })
        ) {
            val isResubmit = it.arguments?.getBoolean("reSubmit", false)

            UploadGovtIssuedId(
                isResubmit = isResubmit,
                onGoBack = {
                    navHostController.popBackStack()
                },
                viewModel = onboardingViewModel,
                onNext = {
                    onNext(it)
                },
            )
        }
        composable(
            route = UploadDisplayPictureDestination.routeWithArgs,
            arguments = listOf(navArgument(UploadDisplayPictureDestination.RESUBMIT_ID_ARG) {
                type = NavType.BoolType
            })
        ) {
            val reSubmit = it.arguments?.getBoolean("reSubmit", false)

            UploadDisplayPicture(
                onGoBack = {
                    navHostController.popBackStack()
                },
                viewModel = onboardingViewModel,
                onNext = {
                    onNext(it)
                },
                reSubmit = reSubmit,
            )
        }
        composable(
            route = MpesaScreenDestination.routeWithArgs,
            arguments = listOf(navArgument(MpesaScreenDestination.ONBOARDING_ID_ARG) {
                type = NavType.StringType
            })
        ) {
            val mpesaViewModel: MpesaViewModel = koinNavViewModel()
            MpesaScreen(
                onGoBack = {
                    navHostController.popBackStack()
                },
                onSuccess = {
                    onNext(SuccessScreenDestination.route)
                },
                viewModel = mpesaViewModel,
                mainViewModel = mainViewModel,
            )
        }
        composable(route = SuccessScreenDestination.route) {
            SuccessScreen(
                onNavigateTo = onNavigateTo,
            )
        }
        composable(route = AccountScreenDestination.route) {
            AccountScreen(
                mainViewModel = mainViewModel,
                onNext = {
                    onNext(it)
                },
                onGoBack = {
                    navHostController.popBackStack()
                }
            )
        }
        composable(route = CreateLandScreenDestination.route) {
            CreateLandScreen(viewModel = onboardingViewModel,
                currentDestination = it.destination,
                onNavigateTo = onNavigateTo,
                mainViewModel = mainViewModel,
                onNext = {
                    onNext(it)
                },
                onClickAddLand = {
                    onNext(OnboardingScreenDestination.route)
                })
        }
        composable(route = SearchScreenDestination.route) {
            SearchLandScreen(currentDestination = it.destination,
                onNavigateTo = onNavigateTo,
                onNavigateToFoundLand = { navHostController.navigate(it) },
                landTitleViewModel = landTitleViewModel,
            )
        }
        composable(route = FoundLandScreenDestination.route) {
            FoundLandScreen(
                onGoBack = { navHostController.popBackStack() },
                viewModel = landTitleViewModel,
            )
        }
        composable(route = OnboardingScreenDestination.route) {
            OnboardingScreen(
                onGoBack = {
                    navHostController.popBackStack()
                },
                onNext = { onNext(it) },
                onboardingViewModel = onboardingViewModel,
            )
        }
    }
}