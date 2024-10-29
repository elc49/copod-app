package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.Badge
import androidx.compose.material3.BadgedBox
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.NavigationBar
import androidx.compose.material3.NavigationBarItem
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.navigation.NavDestination
import androidx.navigation.NavDestination.Companion.hierarchy
import com.lomolo.copodapp.R
import com.lomolo.copodapp.ui.common.Avatar
import com.lomolo.copodapp.ui.common.LandCard
import com.lomolo.copodapp.ui.common.NoLands
import com.lomolo.copodapp.ui.navigation.Navigation
import com.lomolo.copodapp.ui.viewmodels.GetLocalLands
import com.lomolo.copodapp.ui.viewmodels.MarketViewModel
import com.lomolo.copodapp.viewmodels.MainViewModel
import org.koin.androidx.compose.koinViewModel

object ExploreScreenDestination : Navigation {
    override val title = null
    override val route = "explore"
}

sealed class Screen(
    val name: Int,
    val defaultIcon: Int,
    val activeIcon: Int,
    val route: String,
    var showBadge: Boolean = false,
) {
    data object Explore : Screen(
        R.string.explore,
        R.drawable.explore_outlined,
        R.drawable.explore_filled,
        "explore",
        false,
    )
}

data class Land(
    val titleNo: String,
    val town: String,
    val size: Int,
    val symbol: String,
)

val lands = listOf<Land>(
    Land("re/fler/834", "Suneka", 24, "HA"),
    Land("fk/392/rf34", "Ngong", 32000, "HA"),
    Land("df/3892/sdkjd", "Parklands", 48543, "HA"),
    Land("fj/489/283/df", "Upperhill", 983923, "HA")
)

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun MarketScreen(
    modifier: Modifier = Modifier,
    mainViewModel: MainViewModel = koinViewModel<MainViewModel>(),
    marketViewModel: MarketViewModel = koinViewModel<MarketViewModel>(),
    onNavigateTo: (String) -> Unit,
    currentDestination: NavDestination,
) {
    var openDialog by remember { mutableStateOf(false) }
    val navItems = listOf(Screen.Explore)
    val userInfo = mainViewModel.userInfo
    val lands by marketViewModel.lands.collectAsState()

    Scaffold(topBar = {
        TopAppBar(title = {}, actions = {
            Avatar(
                avatar = userInfo.profileImage,
                email = userInfo.email,
                onClick = { openDialog = true },
            )
        })
    }, bottomBar = {
        NavigationBar {
            navItems.forEachIndexed { _, item ->
                // TODO read from current destination
                val isActive = currentDestination.hierarchy.any { it.route == item.route } == true

                NavigationBarItem(selected = false,
                    onClick = { if (!isActive) onNavigateTo(item.route) },
                    icon = {
                        if (item.showBadge) {
                            BadgedBox(badge = { Badge() }) {
                                Icon(
                                    painterResource(if (isActive) item.activeIcon else item.defaultIcon),
                                    modifier = Modifier.size(28.dp),
                                    contentDescription = null,
                                )
                            }
                        } else {
                            Icon(
                                painterResource(if (isActive) item.activeIcon else item.defaultIcon),
                                modifier = Modifier.size(28.dp),
                                contentDescription = null,
                            )
                        }
                    },
                    label = {
                        Text(
                            stringResource(item.name),
                            fontWeight = if (isActive) FontWeight.ExtraBold else FontWeight.Normal,
                        )
                    })
            }
        }
    }) { innerPadding ->
        Surface(
            modifier = modifier.padding(innerPadding)
        ) {
            if (openDialog) {
                AccountDetails(
                    setDialog = { openDialog = it },
                    userInfo = mainViewModel.userInfo,
                    signOut = { mainViewModel.logOut() },
                )
            }
            when (marketViewModel.gettingLands) {
                GetLocalLands.Success -> {
                    if (lands.isEmpty()) {
                        NoLands()
                    } else {
                        LazyColumn(
                            verticalArrangement = Arrangement.spacedBy(4.dp),
                        ) {
                            items(lands) {
                                LandCard(land = it)
                            }
                        }
                    }
                }

                GetLocalLands.Loading -> Column(
                    Modifier.fillMaxSize(),
                    verticalArrangement = Arrangement.Center,
                    horizontalAlignment = Alignment.CenterHorizontally,
                ) {
                    CircularProgressIndicator(
                        Modifier.size(20.dp)
                    )
                }

                is GetLocalLands.Error -> Text(
                    stringResource(R.string.something_wrong),
                    color = MaterialTheme.colorScheme.error,
                    fontWeight = FontWeight.Bold,
                )
            }
        }
    }
}