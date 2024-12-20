package com.lomolo.copodapp.ui.common

import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun TopBar(
    modifier: Modifier = Modifier,
    title: @Composable (() -> Unit) = {},
    navigationIcon: @Composable (() -> Unit) = {}
) {
    TopAppBar(modifier = modifier, title = title, navigationIcon = navigationIcon)
}