package com.lomolo.copodv2.ui.screens

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.lomolo.copodv2.R
import com.lomolo.copodv2.ui.navigation.Navigation
import com.lomolo.copodv2.ui.theme.CopodV2Theme

object HomeScreenDestination : Navigation {
    override val title = null
    override val route = "home"
}

@Composable
fun HomeScreen(
    modifier: Modifier = Modifier,
    onGoToLogin: () -> Unit,
) {
    Box(
        modifier = modifier
            .fillMaxSize()
            .background()
    ) {
        Button(
            onClick = onGoToLogin,
            modifier = Modifier
                .fillMaxWidth()
                .align(Alignment.BottomCenter),
            shape = MaterialTheme.shapes.extraSmall,
            contentPadding = PaddingValues(16.dp),
        ) {
            Text(
                stringResource(R.string.get_started),
                style = MaterialTheme.typography.titleMedium,
            )
        }
    }
}

@Preview
@Composable
fun HomeScreenPreview() {
    CopodV2Theme {
        Box(
            modifier = Modifier.fillMaxSize()
        ) {
            Button(
                onClick = {},
                modifier = Modifier
                    .fillMaxWidth()
                    .align(Alignment.BottomCenter),
                shape = MaterialTheme.shapes.extraSmall,
                contentPadding = PaddingValues(16.dp),
            ) {
                Text(
                    stringResource(R.string.get_started),
                    style = MaterialTheme.typography.titleMedium,
                )
            }
        }
    }
}