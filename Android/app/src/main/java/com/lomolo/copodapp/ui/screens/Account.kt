package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.navigation.NavDestination
import com.lomolo.copodapp.R
import com.lomolo.copodapp.ui.common.BottomNavBar
import com.lomolo.copodapp.ui.common.TopBar
import com.lomolo.copodapp.ui.navigation.Navigation

object AccountScreenDestination: Navigation {
    override val title = null
    override val route = "account"
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun AccountScreen(
    modifier: Modifier = Modifier,
    currentDestination: NavDestination,
    onNavigateTo: (String) -> Unit,
) {
    Scaffold(
        topBar = {
            TopBar(
                title = {
                    Text(stringResource(R.string.account))
                }
            )
        },
        bottomBar = {
            BottomNavBar(currentDestination = currentDestination, onNavigateTo = onNavigateTo)
        }
    ) { innerPadding ->
        Surface(
            modifier = modifier
                .fillMaxSize()
        ) {
            Column(
                Modifier
                    .fillMaxSize()
                    .padding(innerPadding)
            ) {
                Text("Account")
            }
        }
    }
}