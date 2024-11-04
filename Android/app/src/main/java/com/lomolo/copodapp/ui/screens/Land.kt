package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.twotone.Add
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedIconButton
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.setValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.navigation.NavDestination
import com.lomolo.copodapp.R
import com.lomolo.copodapp.ui.common.BottomNavBar
import com.lomolo.copodapp.ui.common.TopBar
import com.lomolo.copodapp.ui.navigation.Navigation
import com.lomolo.copodapp.ui.viewmodels.GetUserLands
import com.lomolo.copodapp.ui.viewmodels.LandViewModel
import com.lomolo.copodapp.ui.viewmodels.MainViewModel
import com.web3auth.core.types.UserInfo
import org.koin.androidx.compose.koinViewModel

object LandScreenDestination: Navigation {
    override val title = null
    override val route = "land"
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun LandScreen(
    modifier: Modifier = Modifier,
    currentDestination: NavDestination,
    onNavigateTo: (String) -> Unit,
    viewModel: LandViewModel = koinViewModel<LandViewModel>(),
    mainViewModel: MainViewModel,
    userInfo: UserInfo?,
    onClickAddLand: () -> Unit,
) {
    LaunchedEffect(Unit) {
        viewModel.getUserLands(userInfo?.email!!)
    }
    val lands by viewModel.lands.collectAsState()
    var openDialog by remember { mutableStateOf(false) }

    Scaffold(
        topBar = {
            TopBar(
                userInfo = userInfo!!,
                onOpenDialog = { openDialog = true },
            )
        },
        bottomBar = {
            BottomNavBar(
                currentDestination = currentDestination,
                onNavigateTo = onNavigateTo,
            )
        }
    ) { innerPadding ->
        Surface(
            modifier = modifier.padding(innerPadding)
        ) {
            if (openDialog) {
                AccountDetails(
                    setDialog = { openDialog = it },
                    userInfo = mainViewModel.userInfo!!,
                    signOut = { mainViewModel.logOut() },
                )
            }
            when(viewModel.gettingUserLands) {
                GetUserLands.Success -> {
                    if (lands.isEmpty()) {
                        NoLand(
                            onClickAddLand = onClickAddLand,
                        )
                    }
                }
                GetUserLands.Loading -> Column(
                    Modifier.fillMaxSize(),
                    verticalArrangement = Arrangement.Center,
                    horizontalAlignment = Alignment.CenterHorizontally,
                ) {
                    CircularProgressIndicator(
                        Modifier.size(20.dp)
                    )
                }
                is GetUserLands.Error -> Text(
                    "Something went wrong",
                    color = MaterialTheme.colorScheme.error,
                    fontWeight = FontWeight.Bold,
                )
            }
        }
    }
}

@Composable
private fun NoLand(
    modifier: Modifier = Modifier,
    onClickAddLand: () -> Unit,
) {
    Column(
        modifier.fillMaxSize(),
        verticalArrangement = Arrangement.Center,
        horizontalAlignment = Alignment.CenterHorizontally,
    ) {
        Icon(
            painterResource(R.drawable.sealed),
            modifier = Modifier.size(60.dp),
            contentDescription = stringResource(R.string.land),
        )
        Text(stringResource(R.string.no_user_lands))
        OutlinedIconButton(
            onClick = onClickAddLand
        ) {
            Icon(
                Icons.TwoTone.Add,
                contentDescription = stringResource(R.string.add),
            )
        }
    }
}