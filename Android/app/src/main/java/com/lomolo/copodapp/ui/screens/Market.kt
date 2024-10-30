package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import com.lomolo.copodapp.ui.navigation.Navigation

object MarketScreenDestination: Navigation {
    override val title = null
    override val route = "explore"
}

@Composable
fun ExploreMarketsScreen(
    modifier: Modifier = Modifier
) {
    Column(
        modifier = modifier.fillMaxSize()
    ) {
        Text("Explore market")
    }
}