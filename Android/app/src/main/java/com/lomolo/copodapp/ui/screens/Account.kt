package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.MoreVert
import androidx.compose.material3.DropdownMenu
import androidx.compose.material3.DropdownMenuItem
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.navigation.NavDestination
import coil.compose.AsyncImage
import coil.request.ImageRequest
import com.lomolo.copodapp.R
import com.lomolo.copodapp.state.viewmodels.MainViewModel
import com.lomolo.copodapp.ui.common.BottomNavBar
import com.lomolo.copodapp.ui.common.TopBar
import com.lomolo.copodapp.ui.navigation.Navigation

object AccountScreenDestination : Navigation {
    override val title = null
    override val route = "account"
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun AccountScreen(
    modifier: Modifier = Modifier,
    currentDestination: NavDestination,
    mainViewModel: MainViewModel,
    onNavigateTo: (String) -> Unit,
) {
    val userInfo = mainViewModel.userInfo
    var expanded by remember { mutableStateOf(false) }

    Scaffold(topBar = {
        TopBar(title = {
            Text(stringResource(R.string.account))
        }, actions = {
            IconButton(onClick = { expanded = true }) {
                Icon(
                    Icons.Default.MoreVert,
                    contentDescription = stringResource(R.string.vertical_dots)
                )
            }
            DropdownMenu(expanded = expanded, onDismissRequest = { expanded = false }) {
                DropdownMenuItem(text = {
                    Text(
                        stringResource(R.string.sign_out),
                        style = MaterialTheme.typography.titleMedium,
                    )
                }, onClick = { expanded = false; mainViewModel.logOut() }, leadingIcon = {
                    Icon(
                        painterResource(R.drawable.logout),
                        modifier = Modifier.size(28.dp),
                        contentDescription = stringResource(R.string.sign_out),
                    )
                })
            }
        })
    }, bottomBar = {
        BottomNavBar(currentDestination = currentDestination, onNavigateTo = onNavigateTo)
    }) { innerPadding ->
        Surface(
            modifier = modifier.fillMaxSize()
        ) {
            Column(
                Modifier
                    .fillMaxSize()
                    .padding(innerPadding),
                verticalArrangement = Arrangement.spacedBy(12.dp),
                horizontalAlignment = Alignment.CenterHorizontally,
            ) {
                Box {
                    Column {
                        AsyncImage(
                            model = ImageRequest.Builder(LocalContext.current)
                                .data(userInfo?.profileImage).crossfade(true).build(),
                            modifier = Modifier
                                .size(120.dp)
                                .clip(CircleShape),
                            placeholder = painterResource(R.drawable.loading_img),
                            error = painterResource(R.drawable.ic_broken_image),
                            contentScale = ContentScale.Crop,
                            contentDescription = stringResource(R.string.account),
                        )
                    }
                }
                Box {
                    Column(
                        horizontalAlignment = Alignment.CenterHorizontally,
                    ) {
                        Text(
                            "${userInfo?.name}",
                            style = MaterialTheme.typography.titleLarge,
                            textAlign = TextAlign.Center,
                        )
                        Text(
                            "${userInfo?.email}",
                            textAlign = TextAlign.Center,
                        )
                    }
                }
            }
        }
    }
}