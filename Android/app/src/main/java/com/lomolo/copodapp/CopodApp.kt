package com.lomolo.copodapp

import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.navigation.NavHostController
import com.lomolo.copodapp.ui.navigation.NavigationHost
import com.lomolo.copodapp.ui.viewmodels.MainViewModel

@Composable
fun CopodApplication(
    modifier: Modifier = Modifier,
    navHostController: NavHostController,
    mainViewModel: MainViewModel,
) {
    NavigationHost(
        modifier = modifier,
        navHostController,
        mainViewModel = mainViewModel,
    )
}