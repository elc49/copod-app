package com.lomolo.copodv2

import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.navigation.NavHostController
import com.lomolo.copodv2.ui.navigation.NavigationHost
import com.lomolo.copodv2.viewmodels.MainViewModel

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