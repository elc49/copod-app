package com.lomolo.copodv2.ui.navigation

import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.navigation.NavHostController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import com.lomolo.copodv2.ui.screens.DashboardScreen
import com.lomolo.copodv2.ui.screens.DashboardScreenDestination
import com.lomolo.copodv2.ui.screens.LoadingScreen
import com.lomolo.copodv2.ui.screens.LoadingScreenDestination
import com.lomolo.copodv2.ui.screens.LoginScreen
import com.lomolo.copodv2.ui.screens.LoginScreenDestination
import com.lomolo.copodv2.ui.screens.Web3SdkErrorScreen
import com.lomolo.copodv2.ui.screens.Web3SdkErrorScreenDestination
import com.lomolo.copodv2.viewmodels.InitializeSdk
import com.lomolo.copodv2.viewmodels.MainViewModel

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
        InitializeSdk.Loading -> LoadingScreenDestination.route
        InitializeSdk.Success -> {
            if (isLoggedIn) {
                DashboardScreenDestination.route
            } else {
                LoginScreenDestination.route
            }
        }

        else -> Web3SdkErrorScreenDestination.route
    }

    NavHost(
        navController = navHostController,
        startDestination = startRoute.toString(),
        route = RootNavigation.route,
    ) {
        composable(route = LoginScreenDestination.route) {
            LoginScreen(
                modifier = modifier,
                mainViewModel = mainViewModel,
            )
        }
        composable(route = DashboardScreenDestination.route) {
            DashboardScreen(
                modifier = modifier,
                mainViewModel = mainViewModel,
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
    }
}