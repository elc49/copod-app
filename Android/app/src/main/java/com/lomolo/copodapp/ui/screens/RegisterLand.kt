package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import com.lomolo.copodapp.R
import com.lomolo.copodapp.ui.navigation.Navigation
import com.lomolo.copodapp.ui.theme.CopodAppTheme

object RegisterLandScreenDestination: Navigation {
    override val title = R.string.register_land
    override val route = "register-land"
}

@Composable
fun RegisterLandScreen(
    modifier: Modifier = Modifier,
) {}

@Preview(showBackground = true)
@Composable
fun RegisterLandPreview() {
    CopodAppTheme {
        Scaffold { innerPadding ->
            Column(
                Modifier.padding(innerPadding)
            ) {
                Text("Register land")
            }
        }
    }
}