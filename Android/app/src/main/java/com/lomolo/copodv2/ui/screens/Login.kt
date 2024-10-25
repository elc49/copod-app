package com.lomolo.copodv2.ui.screens

import androidx.compose.foundation.layout.Column
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import com.lomolo.copodv2.ui.navigation.Navigation
import com.lomolo.copodv2.viewmodels.MainViewModel

object LoginScreenDestination: Navigation {
    override val title = null
    override val route = "login"
}

@Composable
fun LoginScreen(
    modifier: Modifier = Modifier,
    mainViewModel: MainViewModel,
) {
    Column(
        modifier = modifier,
    ) {
        Text("Login")
        Button(
            onClick = { mainViewModel.login() }
        ) {
            Text("Log In")
        }
    }
}