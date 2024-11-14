package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.twotone.CheckCircle
import androidx.compose.material.icons.twotone.Close
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.lomolo.copodapp.R
import com.lomolo.copodapp.ui.common.TopBar
import com.lomolo.copodapp.ui.navigation.Navigation

object SuccessScreenDestination : Navigation {
    override val title = null
    override val route = "success"
}

@Composable
fun SuccessScreen(
    modifier: Modifier = Modifier,
    onNavigateTo: (String) -> Unit,
) {
    Scaffold(topBar = {
        TopBar(navigationIcon = {
            IconButton(
                onClick = {
                    onNavigateTo(LandScreenDestination.route)
                },
            ) {
                Icon(
                    Icons.TwoTone.Close,
                    contentDescription = stringResource(R.string.close),
                )
            }
        })
    }) { innerPadding ->
        Surface(
            modifier.fillMaxSize(),
        ) {
            Column(
                Modifier
                    .fillMaxSize()
                    .padding(innerPadding),
                verticalArrangement = Arrangement.Center,
                horizontalAlignment = Alignment.CenterHorizontally,
            ) {
                Icon(
                    Icons.TwoTone.CheckCircle,
                    contentDescription = stringResource(R.string.ok),
                    modifier = Modifier.size(60.dp),
                    tint = MaterialTheme.colorScheme.primary,
                )
            }
        }
    }
}