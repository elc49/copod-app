package com.lomolo.copodapp

import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.navigation.NavHostController
import com.lomolo.copodapp.state.viewmodels.LandTitleDetailsViewModel
import com.lomolo.copodapp.ui.navigation.NavigationHost
import com.lomolo.copodapp.state.viewmodels.MainViewModel
import com.lomolo.copodapp.state.viewmodels.OnboardingViewModel

@Composable
fun CopodApplication(
    modifier: Modifier = Modifier,
    navHostController: NavHostController,
    mainViewModel: MainViewModel,
    onboardingViewModel: OnboardingViewModel,
    landTitleViewModel: LandTitleDetailsViewModel,
) {
    NavigationHost(
        modifier = modifier,
        navHostController,
        mainViewModel = mainViewModel,
        onboardingViewModel = onboardingViewModel,
        landTitleViewModel = landTitleViewModel,
    )
}